import json
from selenium import webdriver
from scrapy.http import HtmlResponse
from scrapy.utils.python import to_bytes
from selenium.webdriver.common.by import By
from selenium.webdriver.support.ui import WebDriverWait
from selenium.webdriver.support import expected_conditions as EC
from selenium.common.exceptions import TimeoutException


import time


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


URL = "https://www.linkedin.com/jobs/search?trk=guest_homepage-basic_guest_nav_menu_jobs&position=1&pageNum=0&location=United%20States"
# dates=["Any time","Past week","Past 24 hours"]
dates = ["Past 24 hours"]
experience_level = ["Internship", "Entry level", "Mid-Senior", "Director", "Associate"]
salary = ["40,000", "60,000", "80,000", "100,000", "120,000"]
location = state_abbreviations = [
    "AL",
    "AK",
    "AZ",
    "AR",
    "CA",
    "CO",
    "CT",
    "DE",
    "FL",
    "GA",
    "HI",
    "ID",
    "IL",
    "IN",
    "IA",
    "KS",
    "KY",
    "LA",
    "ME",
    "MD",
    "MA",
    "MI",
    "MN",
    "MS",
    "MO",
    "MT",
    "NE",
    "NV",
    "NH",
    "NJ",
    "NM",
    "NY",
    "NC",
    "ND",
    "OH",
    "OK",
    "OR",
    "PA",
    "RI",
    "SC",
    "SD",
    "TN",
    "TX",
    "UT",
    "VT",
    "VA",
    "WA",
    "WV",
    "WI",
    "WY",
]
tech_job_titles = ["Software Engineer"] or [
    "Software Engineer",
    "Senior Software Engineer",
    "Junior Software Engineer",
    "Software Developer",
    "Backend Engineer",
    "Frontend Engineer",
    "Full Stack Engineer",
    "Data Scientist",
    "Senior Data Scientist",
    "Junior Data Scientist",
    "Machine Learning Engineer",
    "Data Analyst",
    "Data Engineer",
    "DevOps Engineer",
    "Senior DevOps Engineer",
    "Cloud Engineer",
    "Cybersecurity Analyst",
    "Information Security Analyst",
    "Security Engineer",
    "Cloud Architect",
    "Solutions Architect",
    "Cloud Consultant",
    "Full Stack Developer",
    "Frontend Developer",
    "Backend Developer",
    "Web Developer",
    "Product Manager",
    "Technical Product Manager",
    "Product Owner",
    "Systems Administrator",
    "IT Administrator",
    "Network Administrator",
    "Database Administrator",
    "Database Engineer",
    "SQL Developer",
]

# TODO: hirarchy to search
# - job_title - time - job_type - experience_level - location - salary
# - job_title - time
browser = load_url(URL)

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
        # browser.find_element(By.XPATH,"//button[@class='filter__submit-button']").click()
        element = browser.find_element_by_class_name("filter__submit-button")
        browser.execute_script("arguments[0].click();", element)

    last_link = "blank"
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
    file_path = "data.json"
    links = response.css(
        ".jobs-search__results-list li .base-card__full-link::attr(href)"
    ).extract()
    storage = []
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
        storage.append(
            {
                "link": link.strip(),
                "job_title": job_title.strip(),
                "company_name": company_name.strip(),
                "address": address.strip(),
                "date_posted": date_posted.strip() if date_posted else None,
                "exact_date": exact_date.strip() if exact_date else None,
            }
        )
        # Specify the file path

        # Write JSON data to a file
    with open(file_path, "w") as json_file:
        for item in storage:
            json.dump(item, json_file, indent=4)  # indent=4 for pretty printing

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
