export let productAddresses = { get: "product/productSrv/GetProductById",
    getAll: "product/productSrv/GetProducts",
    update: "product/productSrv/UpdateProduct",
    create: "product/productSrv/CreateProduct",
    delete: "product/productSrv/DeleteProduct",
    new: "/product/new",
    reload: "product/",
    previousPage: "/product",
    searchPreviousPage: "/"}

export let customerAddresses = { get: "customer/customerSrv/GetCustomerById",
    getAll: "customer/customerSrv/GetCustomers",
    update: "customer/customerSrv/UpdateCustomer",
    create: "customer/customerSrv/CreateCustomer",
    delete: "customer/customerSrv/DeleteCustomer",
    new: "/customer/new",
    reload: "customer/",
    previousPage: "/customer",
    searchPreviousPage: "/"}