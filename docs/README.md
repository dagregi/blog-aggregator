# API Documentation

## Overview

This API serves as an aggregator for RSS feeds. It allows users to save RSS feeds, and the
server periodically scrapes and saves the posts in the feeds storing them in a database.

## Endpoints

The API consists of the following endpoints:

-   [users](https://github.com/dagregi/blog-aggregator/blob/master/docs/users.md)
-   [feeds](https://github.com/dagregi/blog-aggregator/blob/master/docs/feeds.md)
-   [posts](https://github.com/dagregi/blog-aggregator/blob/master/docs/posts.md)
-   [feed_follows](https://github.com/dagregi/blog-aggregator/blob/master/docs/feed_follows.md)

## Authentication

-   Method: API Key
-   Token: A random SHA256 encoded string generated when a user is created
-   Header: Authorization: ApiKey <TOKEN>
