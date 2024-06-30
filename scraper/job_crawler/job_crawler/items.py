# Define here the models for your scraped items
#
# See documentation in:
# https://docs.scrapy.org/en/latest/topics/items.html

import scrapy


class JobCrawlerItem(scrapy.Item):
    # define the fields for your item here like:
    # name = scrapy.Field()
    job_link=scrapy.Field()
    job_location=scrapy.Field()
    job_posted=scrapy.Field()
    exact_date=scrapy.Field()
    company_name=scrapy.Field()
    job_title=scrapy.Field()
    company_link=scrapy.Field()
    job_type=scrapy.Field()
    company_domain=scrapy.Field()
    job_level=scrapy.Field()
