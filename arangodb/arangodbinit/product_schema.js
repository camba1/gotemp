db._createDatabase('product', {},  [{ username: "productUser", passwd: "TestDB@home2", active: true}]);
db._useDatabase('product');
db._create('product');
db._createEdgeCollection("isparentof");
db._createEdgeCollection("iscomponentof");



db.product.insert({"_key": "switch", "name": "Play Switch Console", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "hierarchylevel": "sku", "externalId": "12345"});
db.product.insert({"_key": "tele", "name": "Watch me TV", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "hierarchylevel": "sku"});
db.product.insert({"_key": "fridge", "name": "Cool Stuff Fridge", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "hierarchylevel": "sku"});
db.product.insert({"_key": "intamd", "name": "IntAmd processor", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "hierarchylevel": "sku"});
db.product.insert({"_key": "elec", "name": "Electronics", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "hierarchylevel": "cat"});
db.product.insert({"_key": "appli", "name": "Appliances", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "hierarchylevel": "cat"});

db.isparentof.insert({"_from":"product/elec" , "_to":"product/tele",  "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.isparentof.insert({"_from":"product/elec" , "_to":"product/switch", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.isparentof.insert({"_from":"product/elec" , "_to":"product/intamd", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.isparentof.insert({"_from":"product/appli" , "_to":"product/fridge", "validFrom": "2019-01-02", "validThru": "2021-01-02"});


db.iscomponentof.insert({"_from":"product/intamd" , "_to":"product/switch", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.iscomponentof.insert({"_from":"product/intamd" , "_to":"product/tele", "validFrom": "2019-01-02", "validThru": "2021-01-02"});