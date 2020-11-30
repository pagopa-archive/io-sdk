const RULES_SEPARATOR = '|'
const RULE_PARAMETER_SEPARATOR = ':'

const RULE_NUMERIC = 'numeric'
const RULE_REQUIRED = 'required'
const RULE_MAX_LENGTH = 'max'
const RULE_MIN_LENGTH = 'min'
const RULE_FISCALCODE = 'fiscalCode'

const FIELD_VALUE_KEY = 'value'
const FIELD_RULES_KEY = 'rules'

/**
 * Validates each form's field against a set of rule
 * The field object should contain a value and a rules key, false returne otherwise
 * 
 * @param {array} fields 
 * 
 */
const validateForm = ( fields ) => {

    let found = false;

    Object.keys(fields).forEach((key, index) => {

        const res = fields[key][FIELD_VALUE_KEY] 
            && fields[key][FIELD_RULES_KEY] 
            && validateField( 
                fields[key][FIELD_VALUE_KEY], 
                fields[key][FIELD_RULES_KEY] 
            )

        if(!res) {
            found = true;
        }
            
    })

    return !found

}

/**
 * Validate a single value against a set of rules
 * A set of rules is a "|" separated values
 * 
 * @param  { string } value
 * @param  { string } rules
 */
const validateField = ( value, rules ) => {

    const rulesArray = explodeRules( rules );

    return rulesArray.every( rule => validateFieldRule( value, rule ))

}

const explodeRules = ( rules ) => {

    const rulesArray = rules.split(RULES_SEPARATOR)

    return rulesArray.map( rule => rule.split(RULE_PARAMETER_SEPARATOR))

}

const validateFieldRule = ( field, rule ) => {

    const ruleName = rule[0]
    const ruleParam = rule[1]

    switch( ruleName ) {

        case(RULE_NUMERIC):
            return validateNumericField( field )
        case(RULE_FISCALCODE):
            return validateFiscalCodeField( field )
        case(RULE_MAX_LENGTH):
            return validateMaxLengthField( field, ruleParam )
        case(RULE_MIN_LENGTH):
            return validateMinLengthField( field, ruleParam )
        case(RULE_REQUIRED):
            return validateRequiredField( field )
        default:
            return false
            
    }

}

const validateNumericField = ( value ) => {

    return !isNaN( value )

}

const validateMaxLengthField = ( value, length ) => {

    return value.length <= length

}

const validateMinLengthField = ( value, length ) => {

    return value.length >= length

}

const validateRequiredField = ( value ) => {

    return value.length > 0

}

const validateFiscalCodeField = ( value ) => {

    const fiscalCodePattern = /^[a-zA-Z]{6}[0-9]{2}[a-zA-Z][0-9]{2}[a-zA-Z][0-9]{3}[a-zA-Z]$/

    return value.search(fiscalCodePattern) !== -1

}

export const formValidator = { validateField, validateForm }    