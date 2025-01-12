# Define your item pipelines here
#
# Don't forget to add your pipeline to the ITEM_PIPELINES setting
# See: https://docs.scrapy.org/en/latest/topics/item-pipeline.html


# useful for handling different item types with a single interface
# type: ignore
import psycopg2
from itemadapter import ItemAdapter
from datetime import datetime
import os
import json


class JobCrawlerPipeline:

    def __init__(self) -> None:
        self.items = []
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
            """CREATE TABLE IF NOT EXISTS {} (
           id serial PRIMARY KEY, 
           job_title text,
           job_link text,
           job_location text,
           job_posted text,
           company_name text,
           exact_date text,
           job_type text DEFAULT NULL, 
           company_domain text DEFAULT NULL,
           job_level text DEFAULT NULL,
           company_link text DEFAULT NULL,
           job_description text,
           applicants text,
           UNIQUE (job_title, company_name, job_level))""".format(
                TABLE_NAME
            )
        )

    def process_item(self, item, spider):
        try:
            self.cur.execute(
                """insert into jobs_db (job_title,job_link,job_location,job_posted,company_name,exact_date,job_type,company_domain,job_level,company_link,job_description) values (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)""",
                (
                    item.get("job_title"),
                    item.get("job_link"),
                    item.get("job_location"),
                    item.get("job_posted"),
                    item.get("company_name"),
                    item.get("exact_date"),
                    item.get("employment_type"),
                    item.get("industries"),
                    item.get("seniority_level"),
                    item.get("company_link"),
                    item.get("description"),
                    item.get("applicants"),
                ),
            )

            ## Execute insert of data into database
            self.connection.commit()
        except psycopg2.errors.UniqueViolation:
            self.connection.rollback()
            print("Duplicate job found, not inserting into database.")
        except:
            self.cur.execute("rollback")
            self.cur.execute(
                """insert into jobs_db (job_title,job_link,job_location,job_posted,company_name,exact_date,job_type,company_domain,job_level,company_link,job_description) values (%s,%s,%s,%s,%s,%s,%s,%s,%s,%s,%s)""",
                (
                    item.get("job_title"),
                    item.get("job_link"),
                    item.get("job_location"),
                    item.get("job_posted"),
                    item.get("company_name"),
                    item.get("exact_date"),
                    item.get("employment_type"),
                    item.get("industries"),
                    item.get("seniority_level"),
                    item.get("company_link"),
                    item.get("description"),
                ),
            )
            self.connection.commit()
        self.items.append(dict(item))
        return item

    def close_spider(self, spider):
        date = datetime.now()
        directory = "crawl-tracker"
        if not os.path.exists(directory):
            os.makedirs(directory)

        # Format the filename with date and time
        file_path = os.path.join(
            directory,
            "linkedin-{}.json".format(
                date.strftime(
                    "%a-%d-%m-%Y-%I-%M%p"
                ).lower()  # Ensure the time format is consistent (AM/PM in lowercase)
            ),
        )

        # Open the file and save the items
        with open(file_path, "w") as json_file:
            json.dump(self.items, json_file, indent=4)

        print(f"All items saved to {file_path}")
        self.cur.close()
        self.connection.close()
