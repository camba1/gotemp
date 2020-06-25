db._createDatabase('customer', {},  [{ username: "customerUser", passwd: "TestDB@home2", active: true}]);
db._useDatabase('customer');
db._create('customer');
db._createEdgeCollection("isparentof");

db.customer.insert({"_key": "ducksrus", "name": "Ducks R Us", "validFrom": "2019-01-02", "validThru": "2021-01-02", "externalId":"12345"});
db.customer.insert({"_key": "ducksrusd", "name": "Ducks R Us Dock", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.customer.insert({"_key": "ducksrusb", "name": "Ducks R Us Bill", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.customer.insert({"_key": "patoloco", "name": "Pato Loco Inc", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.customer.insert({"_key": "canard", "name": "Canard Oui Oui", "validFrom": "2019-01-02", "validThru": "2021-01-02"});

db.isparentof.insert({"_from":"customer/ducksrus" , "_to":"customer/ducksrusd", "Type": "soldto", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.isparentof.insert({"_from":"customer/ducksrus" , "_to":"customer/ducksrusb", "Type": "payto", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
