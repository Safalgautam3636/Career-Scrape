# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
import psycopg2
from itemadapter import ItemAdapter


class JobCrawlerPipeline:

    def __init__(self) -> None:
        DB_HOST = "localhost"
        DB_USER = "safalgautam"
        DB_PASSWORD = "yournewpass"
        DB_NAME = "careerscrape"
        DB_PORT = 5432
        TABLE_NAME = "jobs_db"

        self.connection = psycopg2.connect(
            host=DB_HOST,
            user=DB_USER,
            password=DB_PASSWORD,
            dbname=DB_NAME,
            port=DB_PORT,
        )

        self.cur = self.connection.cursor()
        self.cur.execute(
            """CREATE TABLE IF NOT EXISTS {}(id serial PRIMARY KEY, 
                job_title text,
                job_link text,
                job_location text,
                job_posted text,
                company_name text,
                exact_date text,
                job_type text,
                company_domain text,
                job_level text,
                company_link text)""".format(
                TABLE_NAME
            )
        )

    def process_item(self, item, spider):
        try:
            employment_type = item.get("employment_type") or "$"
            industries = item.get("industries") or "$"
            seniority_level=item.get("seniority_level") or "$"
            company_link=item.get("company_link") or "$"
            self.cur.execute(
                """insert into jobs_db (job_title,job_link,job_location,job_posted,company_name,exact_date,job_type,company_domain,job_level,company_link) values (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)""",
                (
                    item["job_title"],
                    item["job_link"],
                    item["job_location"],
                    item["job_posted"],
                    item["company_name"],
                    item["exact_date"],
                    employment_type,
                    industries,
                    seniority_level,
                    company_link
                ),
            )

            ## Execute insert of data into database
            self.connection.commit()

        except:
            self.cur.execute("rollback")
            self.cur.execute(
                """insert into jobs_db (job_title,job_link,job_location,job_posted,company_name,exact_date,job_type,company_domain,job_level,company_link) values (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)""",
                (
                    item["job_title"],
                    item["job_link"],
                    item["job_location"],
                    item["job_posted"],
                    item["company_name"],
                    item["exact_date"],
                    employment_type,
                    industries,
                    seniority_level,
                    company_link,
                ),
            )
            self.connection.commit()
        return item

    def close_spider(self, spider):
        self.cur.close()
        self.connection.close()
