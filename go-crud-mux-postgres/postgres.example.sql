CREATE TABLE pages (
    id SERIAL PRIMARY KEY,
    page_guid VARCHAR(256) NOT NULL DEFAULT '',
    page_title VARCHAR(256) DEFAULT NULL,
    page_content TEXT,
    page_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (page_guid)
);

INSERT INTO pages (page_guid, page_title, page_content, page_date) 
VALUES ("hello-world", "Hello, World", "I\'m so glad you found this page! It\'s been sitting patiently on the Internet for some time, just waiting for a visitor.", CURRENT_TIMESTAMP);

INSERT INTO pages (page_guid, page_title, page_content, page_date) 
VALUES ("a-new-blog", "A New Blog", "I hope you enjoyed the last blog! Well brace yourself, because my latest blog is even <i>better</i> than the last!", "2015-04-29 02:16:19");

INSERT INTO pages (page_guid, page_title, page_content, page_date) 
VALUES ("lorem-ipsum", "Lorem Ipsum", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas sem tortor, lobortis in posuere sit amet, ornare non eros. Pellentesque vel lorem sed nisl dapibus fringilla. In pretium...", "2015-05-06 04:09:45");

