create table countries(
    id int auto_increment,
    name varchar(255) unique,
     is_active tinyint default 0,
    created_at datetime,
    updated_at datetime,
    primary key(id)
);

create table states (
    id int auto_increment,
    name varchar(255) unique,
    country_id  int not null,
    is_active tinyint default 0,
    created_at datetime,
    updated_at datetime,
    primary key(id),
    foreign key(country_id) references countries(id)
);
create table hospitals(
    id int auto_increment,
    name varchar(255) not null,
    address varchar(500) not null,
    city varchar(255) not null,
    state_id int not null,
    country_id int not null,
    is_active tinyint default 0,
    contact_number varchar(255) not null,
    fax_number varchar(255) default null,
    emergency_contact_number varchar(255) not null,
    email varchar(255) default null,
    created_at datetime,
    updated_at datetime,
    primary key(id),
    foreign key(country_id) references countries(id),
    foreign key(state_id) references states(id)
   
);
create table doctors(
    id int auto_increment,
    firstname varchar(255) not null,
    lastname varchar(255) not null,
    email varchar(255) unique,
    contact_number varchar(255) not null,
    hospital_id int not null,
    address varchar(500),
    is_active tinyint default 0,
    created_at datetime,
    updated_at datetime,
     primary key(id),
    foreign key(hospital_id) references hospitals(id)
);

create table users(
    id int auto_increment,
    child_name varchar(255) not null,
    father_name varchar(255) not null,
    mother_name varchar(255) not null,
    date_of_birth date not null,
    gender tinyint default 0,
    mobile varchar(50) not null,
    email varchar(255) unique,
    hospital_id int not null,
    city varchar(255) not null,
    address varchar(500) not null,
    state_id int not null,
    country_id int not null,
    access_token varchar(500) default null,
    is_active tinyint default 0,
    is_verify tinyint default 0,
    last_login_ip varchar(255),
    last_login_time datetime,
    created_at datetime,
    updated_at datetime,
     primary key(id),
    foreign key(hospital_id) references hospitals(id),
    foreign key(state_id) references states(id),
    foreign key(country_id) references countries(id)
);