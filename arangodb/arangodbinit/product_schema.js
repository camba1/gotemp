db._createDatabase('product', {},  [{ username: "productUser", passwd: "TestDB@home2", active: true}]);
db._useDatabase('product');
db._create('product');
db._createEdgeCollection("isparentof");
db._createEdgeCollection("iscomponentof");

db.product.insert({"_key": "switch", "name": "Play Switch Console", "validfrom": "2019-01-02", "validthru": "2021-01-02", "hierarchylevel": "sku"});
db.product.insert({"_key": "tele", "name": "Watch me TV", "validfrom": "2019-01-02", "validthru": "2021-01-02", "hierarchylevel": "sku"});
db.product.insert({"_key": "fridge", "name": "Cool Stuff Fridge", "validfrom": "2019-01-02", "validthru": "2021-01-02", "hierarchylevel": "sku"});
db.product.insert({"_key": "intamd", "name": "IntAmd processor", "validfrom": "2019-01-02", "validthru": "2021-01-02", "hierarchylevel": "sku"});
db.product.insert({"_key": "elec", "name": "Electronics", "validfrom": "2019-01-02", "validthru": "2021-01-02", "hierarchylevel": "cat"});
db.product.insert({"_key": "appli", "name": "Appliances", "validfrom": "2019-01-02", "validthru": "2021-01-02", "hierarchylevel": "cat"});

db.isparentof.insert({"_from":"product/elec" , "_to":"product/tele",  "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.isparentof.insert({"_from":"product/elec" , "_to":"product/switch", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.isparentof.insert({"_from":"product/elec" , "_to":"product/intamd", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.isparentof.insert({"_from":"product/appli" , "_to":"product/fridge", "validfrom": "2019-01-02", "validthru": "2021-01-02"});


db.iscomponentof.insert({"_from":"product/intamd" , "_to":"product/switch", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.iscomponentof.insert({"_from":"product/intamd" , "_to":"product/tele", "validfrom": "2019-01-02", "validthru": "2021-01-02"});