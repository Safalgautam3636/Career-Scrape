

SPIDER_MODULES = ["job_crawler.spiders"]
NEWSPIDER_MODULE = "job_crawler.spiders"

SCRAPEOPS_API_KEY = "5930687e-5fb6-438d-800f-d30159565156"
SCRAPEOPS_PROXY_ENABLED = True

DOWNLOADER_MIDDLEWARES = {
    "scrapeops_scrapy_proxy_sdk.scrapeops_scrapy_proxy_sdk.ScrapeOpsScrapyProxySdk": 725,
}
ITEM_PIPELINES = {
   "job_crawler.pipelines.JobCrawlerPipeline": 300,
}



REQUEST_FINGERPRINTER_IMPLEMENTATION = "2.7"
TWISTED_REACTOR = "twisted.internet.asyncioreactor.AsyncioSelectorReactor"
FEED_EXPORT_ENCODING = "utf-8"
