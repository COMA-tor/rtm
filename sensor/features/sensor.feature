Feature: Sensor value

    In order to work with a measurement
    As a developper
    I need to be able to read a sensor value

    Scenario: Read measurement from a sensor
        Given there is an empty sensor
        When i read the sensor value
        Then the value should be nil

    Scenario: Define sensor value method
        Given there is an empty sensor
        And there is a callback function that return 1
        When the callback function is defined for the sensor
        And i read the sensor value
        Then the value should be 1

    Scenario: Read measurement from a temperature sensor
        Given there is a "temperature" sensor
        When i read the sensor value
        Then the value should nor be nil