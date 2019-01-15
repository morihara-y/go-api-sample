use comment

CREATE TABLE user_comment (
  comment_id int PRIMARY KEY,
  user_id int,
  comment_date date,
  comment varchar(2000) NOT NULL
);

CREATE TABLE comment_score (
  comment_id int PRIMARY KEY,
  score_type int,
  score float
);
