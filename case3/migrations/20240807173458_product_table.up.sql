BEGIN
;

CREATE TABLE IF NOT EXISTS "products" (
    "id" SERIAL PRIMARY KEY,
    "device_name" VARCHAR(50),
    "manufacturer" VARCHAR(50),
    "price" INT,
    "image" TEXT,
    "stock" INT,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NULL
);

CREATE TRIGGER "log_product_update" BEFORE
UPDATE
    ON "products" FOR EACH ROW EXECUTE PROCEDURE log_update_master();

insert into products (device_name, manufacturer, price, image, stock) values ('Lava A55', 'Lava', 875837, 'http://dummyimage.com/164x257.png/cc0000/ffffff', 25);
insert into products (device_name, manufacturer, price, image, stock) values ('Nokia 2.2', 'Nokia', 160441, 'http://dummyimage.com/246x332.png/ff4444/ffffff', 28);
insert into products (device_name, manufacturer, price, image, stock) values ('Infinix Hot 10', 'Infinix', 504545, 'http://dummyimage.com/202x396.png/5fa2dd/ffffff', 17);
insert into products (device_name, manufacturer, price, image, stock) values ('Lenovo Tab 7', 'Lenovo', 973451, 'http://dummyimage.com/214x293.png/cc0000/ffffff', 3);
insert into products (device_name, manufacturer, price, image, stock) values ('Samsung T100', 'Samsung', 730818, 'http://dummyimage.com/128x477.png/cc0000/ffffff', 4);
insert into products (device_name, manufacturer, price, image, stock) values ('Sony Ericsson K850', 'Sony', 855916, 'http://dummyimage.com/190x340.png/ff4444/ffffff', 16);
insert into products (device_name, manufacturer, price, image, stock) values ('BLU J4', 'BLU', 528839, 'http://dummyimage.com/154x433.png/cc0000/ffffff', 27);
insert into products (device_name, manufacturer, price, image, stock) values ('Vertu Diamond', 'Vertu', 953404, 'http://dummyimage.com/219x412.png/dddddd/000000', 29);
insert into products (device_name, manufacturer, price, image, stock) values ('LG GD880 Mini', 'LG', 268981, 'http://dummyimage.com/119x294.png/5fa2dd/ffffff', 26);
insert into products (device_name, manufacturer, price, image, stock) values ('ZTE Boost Max+', 'ZTE', 953345, 'http://dummyimage.com/200x255.png/5fa2dd/ffffff', 4);
insert into products (device_name, manufacturer, price, image, stock) values ('BlackBerry Q10', 'BlackBerry', 111484, 'http://dummyimage.com/105x360.png/cc0000/ffffff', 0);
insert into products (device_name, manufacturer, price, image, stock) values ('Allview P5 AllDro', 'Allview', 193948, 'http://dummyimage.com/238x248.png/cc0000/ffffff', 25);
insert into products (device_name, manufacturer, price, image, stock) values ('Nokia 7110', 'Nokia', 514416, 'http://dummyimage.com/183x284.png/dddddd/000000', 17);
insert into products (device_name, manufacturer, price, image, stock) values ('ZTE Axon 7 mini', 'ZTE', 409132, 'http://dummyimage.com/191x285.png/5fa2dd/ffffff', 12);
insert into products (device_name, manufacturer, price, image, stock) values ('Micromax Bharat 2 Ultra', 'Micromax', 551409, 'http://dummyimage.com/247x241.png/ff4444/ffffff', 24);
insert into products (device_name, manufacturer, price, image, stock) values ('Energizer Energy 100 (2017)', 'Energizer', 443431, 'http://dummyimage.com/116x488.png/dddddd/000000', 12);
insert into products (device_name, manufacturer, price, image, stock) values ('Gionee Marathon M5 mini', 'Gionee', 339573, 'http://dummyimage.com/196x354.png/cc0000/ffffff', 25);
insert into products (device_name, manufacturer, price, image, stock) values ('Oppo N1 mini', 'Oppo', 490458, 'http://dummyimage.com/215x392.png/dddddd/000000', 7);
insert into products (device_name, manufacturer, price, image, stock) values ('Nokia Asha 306', 'Nokia', 876458, 'http://dummyimage.com/177x317.png/ff4444/ffffff', 26);
insert into products (device_name, manufacturer, price, image, stock) values ('LG Optimus Black (White version)', 'LG', 188822, 'http://dummyimage.com/207x279.png/dddddd/000000', 15);
insert into products (device_name, manufacturer, price, image, stock) values ('Micromax A121 Canvas Elanza 2', 'Micromax', 619476, 'http://dummyimage.com/185x232.png/ff4444/ffffff', 9);
insert into products (device_name, manufacturer, price, image, stock) values ('Pantech Element', 'Pantech', 912596, 'http://dummyimage.com/149x225.png/dddddd/000000', 2);
insert into products (device_name, manufacturer, price, image, stock) values ('Plum Whiz', 'Plum', 636689, 'http://dummyimage.com/202x282.png/5fa2dd/ffffff', 0);
insert into products (device_name, manufacturer, price, image, stock) values ('Allview P8 Energy mini', 'Allview', 767267, 'http://dummyimage.com/107x460.png/ff4444/ffffff', 1);
insert into products (device_name, manufacturer, price, image, stock) values ('alcatel Pixi 3 (3.5)', 'alcatel', 183495, 'http://dummyimage.com/240x286.png/5fa2dd/ffffff', 13);
insert into products (device_name, manufacturer, price, image, stock) values ('Acer Liquid M330', 'Acer', 449028, 'http://dummyimage.com/208x437.png/5fa2dd/ffffff', 12);
insert into products (device_name, manufacturer, price, image, stock) values ('LG W31', 'LG', 614709, 'http://dummyimage.com/103x453.png/cc0000/ffffff', 21);
insert into products (device_name, manufacturer, price, image, stock) values ('alcatel OT-S218', 'alcatel', 673399, 'http://dummyimage.com/176x292.png/5fa2dd/ffffff', 25);
insert into products (device_name, manufacturer, price, image, stock) values ('Samsung S5330 Wave533', 'Samsung', 941171, 'http://dummyimage.com/156x387.png/ff4444/ffffff', 3);
insert into products (device_name, manufacturer, price, image, stock) values ('O2 XDA Exec', 'O2', 532959, 'http://dummyimage.com/134x226.png/dddddd/000000', 28);

COMMIT;