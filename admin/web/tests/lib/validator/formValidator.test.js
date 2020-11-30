import { formValidator } from '../../../src/lib/validator/formValidator';

describe("Validate field", () => {

    it("Should return true when validating a numeric rule", () => {

        const res = formValidator.validateField( "123", "numeric")

        expect(res).toBe(true)

    });

    it("Should return true when validating a fiscal code", () => {

        const res = formValidator.validateField("SLVNCI09A15H501A", "required|fiscalCode")

        expect(res).toBe(true)

    })

    it("Should return true when validating a max length field", () => {

        const res = formValidator.validateField( "valueMaxLengthOk", "required|max:300")

        expect(res).toBe(true)

    })

    it("Should return true when validating min length field", () => {

        const res = formValidator.validateField("","min:0")

        expect(res).toBe(true)

    })

    it("Should return false on bad fiscalcode", () => {

        const res = formValidator.validateField("SLNCI09A15H501A", "required|fiscalCode")
        
        expect(res).toBe( false )

    })

    it("Should return false on longer max length", () => {

        const res = formValidator.validateField("valueLonger", "required|max:5")
        
        expect(res).toBe( false )

    })

    it("Should return false on shorter than min length", () => {

        const res = formValidator.validateField("valueShorter", "required|min:50")
        
        expect(res).toBe( false )

    })

    it("Should return false on non numeric value", () => {

        const res = formValidator.validateField("123NaN", "numeric")

        expect(res).toBe(false)

    })

})

const cases = [
    [
        {
            data:  {
                numericField : {
                    "name": "numericField",
                    "value": "123",
                    "rules": "required|numeric"
                },
                fiscalCodeField : {
                    "name": "fiscalCodeField",
                    "value": "SLVNCI09A15H501A",
                    "rules": "required|fiscalCode"
                },
                stringField1 : {
                    "name": "stringField1",
                    "value": "someRandomString",
                    "rules": "required|max:255|min:1"
                },
                stringField2: {
                    "name": "stringField2",
                    "value": "someRandomString2",
                    "rules": "required|min:10"
                }
            },
            expectedResult: true
        },
        {
            data:  {
                numericField : {
                    "name": "numericField",
                    "value": "123",
                    "rules": "required|numeric"
                },
                fiscalCodeField : {
                    "name": "fiscalCodeField",
                    "value": "SLVNCI09A15H501A",
                    "rules": "required|fiscalCode"
                },
                stringField1 : {
                    "name": "stringField1",
                    "value": "someRandomString",
                    "rules": "required|max:255|min:1"
                },
                stringField2: {
                    "name": "stringField2",
                    "value": "someRandomString2",
                    "rules": "required|min:100"
                }
            },
            expectedResult: false
        },
        {
            data:  {
                numericField : {
                    "name": "numericField",
                    "value": "123",
                    "rules": "required|numeric"
                },
                fiscalCodeField : {
                    "name": "fiscalCodeField",
                    "value": "SLVNCI09A15H501A",
                    "rules": "required|fiscalCode"
                },
                stringField1 : {
                    "name": "stringField1",
                    "value": "someRandomString",
                    "rules": "required|max:5|min:1"
                },
                stringField2: {
                    "name": "stringField2",
                    "value": "someRandomString2",
                    "rules": "required|min:10"
                }
            },
            expectedResult: false
        },
        {
            data:  {
                numericField : {
                    "name": "numericField",
                    "value": "123",
                    "rules": "required|numeric"
                },
                fiscalCodeField : {
                    "name": "fiscalCodeField",
                    "value": "SLVCI0A15H501A",
                    "rules": "required|fiscalCode"
                },
                stringField1 : {
                    "name": "stringField1",
                    "value": "someRandomString",
                    "rules": "required|max:255|min:1"
                },
                stringField2: {
                    "name": "stringField2",
                    "value": "someRandomString2",
                    "rules": "required|min:10"
                }
            },
            expectedResult: false
        },
        {
            data:  {
                numericField : {
                    "name": "numericField",
                    "value": "123NaN",
                    "rules": "required|numeric"
                },
                fiscalCodeField : {
                    "name": "fiscalCodeField",
                    "value": "SLVNCI09A15H501A",
                    "rules": "required|fiscalCode"
                },
                stringField1 : {
                    "name": "stringField1",
                    "value": "someRandomString",
                    "rules": "required|max:255|min:1"
                },
                stringField2: {
                    "name": "stringField2",
                    "value": "someRandomString2",
                    "rules": "required|min:10"
                }
            },
            expectedResult: false
        },
        {
            data:  {
                numericField : {
                    "name": "numericField",
                    "value": "123",
                    "rules": "required|numeric"
                },
                fiscalCodeField : {
                    "name": "fiscalCodeField",
                    "value": "PZZVLR92R14E783O",
                    "rules": "required|fiscalCode"
                },
                stringField1 : {
                    "name": "stringField1",
                    "value": "someRandomString",
                    "rules": ""
                },
                stringField2: {
                    "name": "stringField2",
                    "value": "someRandomString2",
                    "rules": "required|min:10"
                }
            },
            expectedResult: true
        }
    ]
]

describe('Form Validator', () => {

    it.each(cases)("Should return expected result", ( testCase ) => {

        const res = formValidator.validateForm( testCase.data )

        expect(res).toBe(testCase.expectedResult)

    })

});