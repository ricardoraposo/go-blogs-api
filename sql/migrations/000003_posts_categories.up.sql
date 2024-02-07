CREATE TABLE posts_categories (
    post_id CHAR(36),
    category_id CHAR(36),
    PRIMARY KEY (post_id, category_id),
    FOREIGN KEY (post_id) REFERENCES blog_posts(id),
    FOREIGN KEY (category_id) REFERENCES categories(id)
);
