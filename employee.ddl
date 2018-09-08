CREATE TABLE employee (
  employee_id integer primary key autoincrement not null
  , login_id varchar(255) not null
  , password varchar(255) not null
  , password_lst_chg_dt varchar(10) not null
  , emp_no     varchar(50) not null
  , user_nm_ja varchar(50) not null
  , user_nm_ka varchar(50) not null
  , user_nm_en varchar(50) not null
  , email      varchar(255) not null
  , join_dt    varchar(10) not null
  , quit_dt    varchar(10)
  , version_no int not null
  , created_at timestamp not null
  , created_by varchar(50) not null
  , updated_at timestamp not null
  , updated_by varchar(50) not null
  , is_actived varchar(1) not null
);
