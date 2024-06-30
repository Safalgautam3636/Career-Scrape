import scrapy
import json
import time
from selenium import webdriver
from scrapy.http import HtmlResponse
from scrapy.utils.python import to_bytes
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC

# type Job struct {
# 	gorm.Model
# 	ID              uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
# 	Role            string    `json:"role"`
# 	Link            string    `json:"link"`
# 	Address         string    `json:"address"`
# 	CompanyName     string    `json:"company_name"`
# 	DatePosted      time.Time `json:"posted_date"`
# 	AboutJob        string    `json:"about_job"`
# 	Qualifications  string    `json:"qualifications"`
# 	Responsibility  string    `json:"responsibility"`
# 	Benefits        string    `json:"benefits"`
# 	Salary          string    `json:"salary"`
# 	Yoe             int       `json:"yoe"` //years of experience
# 	PulledTimeStamp time.Time `json:"pulled_date" gorm:"default:CURRENT_TIMESTAMP"`
# 	// CreatedAt       time.Time `json:"created_at,omitempty"`
# 	// UpdatedAt       time.Time `json:"updated_at,omitempty"`
# 	// DeletedAt       time.Time `json:"deleted_at,omitempty"`
# }


def load_url(url):
    browser = webdriver.Chrome(
        executable_path="/Users/safalgautam/Documents/chromedriver-mac-arm64/chromedriver"
    )
    browser.delete_all_cookies()
    browser.get(url)
    return browser


def get_response_from_selenium(browser):
    body = to_bytes(browser.page_source)
    response = HtmlResponse(browser.current_url, body=body, encoding="utf-8")
    return response


class LinkedinJobScraper(scrapy.Spider):
    name = "linkedin_jobs"
    URL = "https://www.linkedin.com/jobs/search?trk=guest_homepage-basic_guest_nav_menu_jobs&position=1&pageNum=0&location=United%20States"
    def start_requests(self):
        # urls = [
        #     "https://quotes.toscrape.com/page/1/",
        #     "https://quotes.toscrape.com/page/2/",
        # ]
        yield scrapy.Request(url=self.URL, callback=self.parse_selenium)
        # for url in urls:
        #     yield scrapy.Request(url=url, callback=self.parse)

    def parse_selenium(self, response):
        # URL = "https://www.linkedin.com/jobs/search?trk=guest_homepage-basic_guest_nav_menu_jobs&position=1&pageNum=0&location=United%20States"
        # dates=["Any time","Past week","Past 24 hours"]
        dates = ["Past 24 hours"]
        tech_job_titles = ["Software Engineer"]
        # TODO: hirarchy to search
        # - job_title - time - job_type - experience_level - location - salary
        # - job_title - time
        browser = load_url(self.URL)

        response = get_response_from_selenium(browser)
        time.sleep(10)

        job_search_field = browser.find_element(
            By.XPATH, "//input[contains(@aria-label,'Search job')]"
        )

        if job_search_field and job_search_field.text != "":
            job_search_field = job_search_field.clear()

        for item in tech_job_titles:
            job_search_field.clear()
            job_search_field.send_keys(item)

            browser.find_element(
                By.XPATH,
                "//input[contains(@value,'public_jobs_jobs-search-bar_search-submit')]/following-sibling::button",
            ).click()

            time.sleep(5)
            response = get_response_from_selenium(browser)
            browser.find_element(By.XPATH, "//ul[@class='filters__list']/li[1]").click()
            time.sleep(5)
            for date in dates:
                time.sleep(5)
                browser.find_element(
                    By.XPATH,
                    "//div[contains(@class,'filter-values-container__filter-value')]/label[contains(text(),'{}')]/..".format(
                        date
                    ),
                ).click()
                element = browser.find_element_by_class_name("filter__submit-button")
                browser.execute_script("arguments[0].click();", element)

            last_link = ""
            current_link = response.css(
                ".jobs-search__results-list li .base-card__full-link::attr(href)"
            ).extract()[-1]
            count = 0
            while (
                current_link == last_link and count == 5
            ) or "You've viewed all jobs for this search" not in str(response.body):
                if current_link == last_link and count == 5:
                    break
                try:
                    response = get_response_from_selenium(browser)
                    if response.xpath('//button[contains(text(),"See more jobs")]'):
                        WebDriverWait(browser, 5).until(
                            EC.element_to_be_clickable(
                                (By.XPATH, '//button[contains(text(),"See more jobs")]')
                            )
                        )
                    browser.find_element(
                        By.XPATH, '//button[contains(text(),"See more jobs")]'
                    ).click()

                    time.sleep(5)
                except:
                    browser.execute_script("window.scrollTo(0, 99999)")
                    if count > 0:
                        browser.execute_script("window.scrollTo(0, 0)")
                        time.sleep(2)
                        browser.execute_script("window.scrollTo(0, 99999)")
                        count = 0
                last_link = current_link
                current_link = response.css(
                    ".jobs-search__results-list li .base-card__full-link::attr(href)"
                ).extract()[-1]
                print("C: ", current_link, "    Last: ", last_link)
                print("Count: ", count)
                if current_link == last_link:
                    count += 1
                else:
                    count = 0
            storage=[]
            for item in response.css(".jobs-search__results-list li"):
                link = item.css(".base-card__full-link::attr(href)").extract_first()
                sub_item = item.css(".base-search-card__info")
                job_title = item.css(".base-search-card__title::text").extract_first()
                company_name = sub_item.css(
                    ".base-search-card__subtitle a::text"
                ).extract_first()
                address = sub_item.css(
                    ".base-search-card__metadata .job-search-card__location::text"
                ).extract_first()
                date_posted = sub_item.css(
                    ".base-search-card__metadata .job-search-card__listdate--new::text"
                ).extract_first()
                exact_date = sub_item.css(
                    ".base-search-card__metadata .job-search-card__listdate--new::attr(datetime)"
                ).extract_first()

                jobs= {
                    "job_link": link.strip() if link else "",
                    "job_title": job_title.strip() if job_title else "",
                    "company_name": company_name.strip() if company_name else "",
                    "job_location": address.strip() if address else "",
                    "job_posted": date_posted.strip() if date_posted else "",
                    "exact_date": exact_date.strip() if exact_date else "",
                }
                storage.append(jobs)
            for job in jobs[:3]:
                yield scrapy.Request(
                    url=job.get("job_link"),
                    callback=self.parse_body,
                    meta={
                        "job_title": job.get("job_title"),
                        "company_name": job.get("company_name"),
                        "job_location": job.get("job_location"),
                        "job_posted": job.get("job_posted"),
                        "exact_date": job.get("exact_date"),
                    },
                )

    def get_info(self, response):
        info = {}
        for item in response.css(".description__job-criteria-list li"):
            k = item.css("h3::text").extract_first()
            v = item.css("span::text").extract_first()
            if k and v:
                info[k.strip().lower()] = v.strip().lower()
        return info

    def parse_body(self, response):
        info = self.get_info(response)
        root_info = response.meta.get("job_des")

        item = {"description": response.css(".show-more-less-html__markup").extract()}
        item.update(info)
        item.update(root_info)
        yield item

    # def parse(self, response):
    #     page_id = int(response.meta.get("id"))
    #     jobs = response.css("body li")
    #     jobs_len = len(jobs)

    #     job_dics = {}
    #     for job in jobs:
    #         job_link = (
    #             job.css(".base-card__full-link::attr(href)").extract_first().strip()
    #             if job.css(".base-card__full-link::attr(href)").extract_first()
    #             else None
    #         )
    #         job_location = (
    #             job.css(".job-search-card__location::text").extract_first().strip()
    #             if job.css(".job-search-card__location::text").extract_first()
    #             else None
    #         )
    #         job_posted = (
    #             job.css("time::text").extract_first().strip()
    #             if job.css("time::text").extract_first()
    #             else None
    #         )
    #         company_name = (
    #             job.css("h4.base-search-card__subtitle a::text").extract_first().strip()
    #             if job.css("h4.base-search-card__subtitle a::text").extract_first()
    #             else None
    #         )
    #         company_link = (
    #             job.css("h4.base-search-card__subtitle a::attr(href)")
    #             .extract_first()
    #             .strip()
    #             if job.css(
    #                 "h4.base-search-card__subtitle a::attr(href)"
    #             ).extract_first()
    #             else None
    #         )
    #         job_title = (
    #             job.css(".base-search-card__title::text").extract_first().strip()
    #             if job.css(".base-search-card__title::text").extract_first()
    #             else None
    #         )

    #         job_dic = {
    #             job_link: {
    #                 "job_link": job_link,
    #                 "job_location": job_location,
    #                 "job_posted": job_posted,
    #                 "company_name": company_name,
    #                 "job_title": job_title,
    #                 "company_link": company_link,
    #             }
    #         }
    #         job_dics.update(job_dic)
    #         # yield job_dic
    #     for link, job_des in job_dics.items():
    #         yield scrapy.Request(
    #             url=link, callback=self.parse_body, meta={"job_des": job_des}
    #         )

    #     if jobs_len > 0:
    #         page_id += 25
    #         job_url = self.job_url + str(page_id)
    #         yield scrapy.Request(url=job_url, callback=self.parse, meta={"id": page_id})

    # Specify the file path

    # Write JSON data to a file

    # with open(file_path, 'w') as json_file:
    #     for item in storage:
    #         json.dump(item, json_file, indent=4)  # indent=4 for pretty printing

    # TODO fiilter by past 24hrs jobs
    # browser.find_element(By.XPATH,"//ul[@class='filters__list']/li[1]").click()
    # browser.find_element(By.XPATH,"//div[contains(@class,'filter-values-container__filter-value')]/label[contains(text(),'Past 24')]/..").click()
    # browser.find_element(By.XPATH,"//button[@class='filter__submit-button']").click()

    # TODO for time
    # element = browser.find_element_by_class_name('filter__submit-button')
    # browser.execute_script("arguments[0].click();", element)

    # TODO location
    # browser.find_element(By.XPATH,"//ul[@class='filters__list']/li[1]").click()

    # browser.find_element(By.XPATH,"//div[contains(@class,'filter-values-container__filter-value')]/label[contains(text(),'NY')]/..").click()
    # element = browser.find_element_by_class_name('filter__submit-button')
    # browser.execute_script("arguments[0].click();", element)
    # pull links response.css(".jobs-search__results-list li .base-card__full-link::attr(href)").extract()

    # TODO for type of job(internship/fulltime)
    # browser.find_element(By.XPATH,"//*[contains(@id,'f_E-0')]").click()
    # element = browser.find_element_by_class_name('filter__submit-button')
    # browser.execute_script("arguments[0].click();", element)
