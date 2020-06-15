db._createDatabase('customer', {},  [{ username: "customerUser", passwd: "TestDB@home2", active: true}]);
db._useDatabase('customer');
db._create('customer');
db._createEdgeCollection("isparentof");

db.customer.insert({"_key": "ducksrus", "name": "Ducks R Us", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.customer.insert({"_key": "ducksrusd", "name": "Ducks R Us Dock", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.customer.insert({"_key": "ducksrusb", "name": "Ducks R Us Bill", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.customer.insert({"_key": "patoloco", "name": "Pato Loco Inc", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.customer.insert({"_key": "canard", "name": "Canard Oui Oui", "validfrom": "2019-01-02", "validthru": "2021-01-02"});

db.isparentof.insert({"_from":"customer/ducksrus" , "_to":"customer/ducksrusd", "Type": "soldto", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
db.isparentof.insert({"_from":"customer/ducksrus" , "_to":"customer/ducksrusb", "Type": "payto", "validfrom": "2019-01-02", "validthru": "2021-01-02"});
