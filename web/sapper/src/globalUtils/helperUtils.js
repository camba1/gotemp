
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