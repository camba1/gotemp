db._createDatabase('customer', {},  [{ username: "customerUser", passwd: "TestDB@home2", active: true}]);
db._useDatabase('customer');
db._create('customer');
db._createEdgeCollection("isparentof");



db.customer.insert({"_key": "ducksrus","name": "Ducks R Us", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "externalId":"12345"});
db.customer.insert({"_key": "ducksrusd","name": "Ducks R Us Dock", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "externalId":"12345"});
db.customer.insert({"_key": "ducksrusb","name": "Ducks R Us Bill", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "externalId":"12345"});
db.customer.insert({"_key": "patoloco","name": "Pato Loco Inc", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "externalId":"12345"});
db.customer.insert({"_key": "canard","name": "Canard Oui Oui", "validityDates": {"validFrom": {"seconds": 1592692598,"nanos": 274583000}, "validThru": {"seconds": 1615792598, "nanos": 274584000}}, "externalId":"12345"});

db.isparentof.insert({"_from":"customer/ducksrus" , "_to":"customer/ducksrusd", "Type": "soldto", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
db.isparentof.insert({"_from":"customer/ducksrus" , "_to":"customer/ducksrusb", "Type": "payto", "validFrom": "2019-01-02", "validThru": "2021-01-02"});
