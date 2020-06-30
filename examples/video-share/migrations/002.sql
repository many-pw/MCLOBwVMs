alter table videos add column worker varchar(255);
alter table videos add column ext varchar(255);
alter table videos add unique index(worker);
