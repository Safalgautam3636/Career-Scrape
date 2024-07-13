# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class JobCrawlerItem(scrapy.Item):
    job_link=scrapy.Field()
    job_location=scrapy.Field()
    job_posted=scrapy.Field()
    exact_date=scrapy.Field()
    company_name=scrapy.Field()
    job_title=scrapy.Field()
    company_link=scrapy.Field()
    employment_type=scrapy.Field()
    industries=scrapy.Field()
    seniority_level=scrapy.Field()
    description=scrapy.Field()
