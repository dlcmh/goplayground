# rake -f db_setup.rb

task default: :db_setup

desc 'Set up development database'
task :db_setup do
  cmd = 'dropdb goplayground'
  p cmd
  `#{cmd}`

  cmd = 'createdb goplayground'
  p cmd
  `#{cmd}`

  cmd = <<~SQL
    psql -d goplayground -c "
      CREATE TABLE pages (
        id BIGSERIAL PRIMARY KEY,
        guid VARCHAR NOT NULL UNIQUE,
        title VARCHAR NOT NULL,
        content TEXT NOT NULL,
        published_on TIMESTAMP WITH TIME ZONE
      )
    "
  SQL
  p cmd
  `#{cmd}`

  content = [
    "I''m so glad you found this page!",
    "It''s been sitting patiently on the Internet for some time,",
    "just waiting for a visitor.",
  ].join(' ')
  cmd = <<~SQL
  psql -d goplayground -c "
    INSERT INTO pages (guid, title, content, published_on)
    VALUES
      ('hello-world', 'Hello, World', '#{content}', NOW());
  "
  SQL
  p cmd
  `#{cmd}`

  content = [
    "I hope you enjoyed the last blog post!",
    "Well brace yourself, because my latest post is even",
    "<i>better</i> than the last!",
  ].join(' ')
  cmd = <<~SQL
  psql -d goplayground -c "
    INSERT INTO pages (guid, title, content, published_on)
    VALUES
      ('a-new-blog', 'A New Blog Post', '#{content}', NOW());
  "
  SQL
  p cmd
  `#{cmd}`
end
