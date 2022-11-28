CREATE DATABASE alif_12; --buat database

--===========================================--

CREATE DATABASE new_db_tools;

use new_db_tools; --buat database

DROP table mentor; -- hapus table

--tambah tambel
CREATE TABLE Mentor( 
    id_Mentor VARCHAR(5) primary key,
    nama VARCHAR (255) not NULL,
    email VARCHAR(255) not NULL,
    alamat VARCHAR(255) not NULL,
    domisili VARCHAR(255) not NULL,
    hp VARCHAR(15) not NULL,
    hobby VARCHAR(255),
    create_at TIMESTAMP default current_timestamp(),
    is_delete boolean default false
);

--data paling belakang tidak perlu menggunakan coma
create table `class`(
    id_class VARCHAR(5),
    nama VARCHAR(255),
    mentor VARCHAR(5),
    start datetime not NULL,
    end datetime not null,
    create_at TIMESTAMP default current_timestamp(),
    primary key (id_class),
    constraint foreign  key(mentor) references mentor(id_Mentor) -- tambah constraint untuk foreign key
); --on delete cascode untuk memisahkan atau membuat jarak

--modif sebuah table
ALTER table `class` drop column end; --menghapus data
ALTER table `class` add datetime not null; --mengembalikan data kembali
ALTER TABLE `class` drop constraint class_ibfk_1; -- cara menghapus foreign key
ALTER TABLE `class` add constraint fk_class_mentor foreign key(mentor) references mentor(id_Mentor)

rename table `class` to kalas; -- cara mengganti nama class ke kelas