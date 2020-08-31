
/**
 * Checks if an object is empty
 * @param obj - Object ot check
 * @returns {boolean} - True if object is empty
 */
export function isObjectEmpty(obj) {
    for(let i in obj) return false;
    return true;
}

/**
 * Checks if object is a valid date
 * @param d - object to check
 * @returns {boolean} - True if object is a valid date
 */
export function isValidDate(d) {
    return d instanceof Date && !isNaN(d);
}

/**
 * Check if a string canbe coverted to a valid date
 * @param stringDate - String possibly containing a date
 * @returns {boolean} - True if string contains a valid date
 */
export function isValidStringDate(stringDate) {
    if (stringDate === "") {
        return false
    }
    let d = new Date(stringDate)
    return isValidDate(d)
}

/**
 * Take an object and convert it to name, value KV pairs
 * @param obj - Object ot convert
 * @returns {{string, string}[]} - KV pairs
 */
export function convertExtraFields(obj) {
    let vals = []
    let key;
    if (!isObjectEmpty(obj)) {
        for (let key1 in obj) {
            key = key1
            vals.push({'name': key, 'value': obj[key]})
        }
        return vals
    }
}

/**
 * Since some fields come as date time and we need to display them as dates,
 * we need to update them manually
 * @param dateToUpdate - String, name of the date field to update
 * @param newDateString - String that contains the date to be used for update
 * @param objectToUpdate - Object to be updated
 */
export function updateValidDate(dateToUpdate, newDateString, objectToUpdate){
    let foundVD = false
    if (objectToUpdate) {
            let parts = newDateString.split('-');
            let VD = new Date(parts[0], parts[1] - 1, parts[2]);
            if (isValidDate(VD)){
                objectToUpdate[dateToUpdate] = VD
            }
            foundVD = true
    }
    if (!foundVD) {
        alert('Unable to populate validity date')
    }
}

/**
 * By default grpc will not send out empty fields, so our object may be missing some fields.
 * Additionally, GO will not unMarshal a string value to a boolean field.
 * So we need to initialize the boolean fields with a boolean value so that it is not
 * automatically marshalled as a string
 * @param objToCheck - Object to which we may need to add the boolean field
 * @param flagName - Name of the boolean field to add
 */
export function addBoolField(objToCheck, flagName) {
    if (objToCheck && !objToCheck.hasOwnProperty(flagName)) {
        objToCheck[flagName] = false
    }
}

/**
 * Check if an object has a property, if it does not add it with initial value = passed in value
 * @param objToCheck - Object to which we may need to add a property
 * @param fieldName - Name of the  field to add
 * @param value - Initial default value for the field to add
 */
export function addFieldtoObj(objToCheck, fieldName, value) {
    if (objToCheck && !objToCheck.hasOwnProperty(fieldName)) {
        objToCheck[fieldName] = value
    }
}