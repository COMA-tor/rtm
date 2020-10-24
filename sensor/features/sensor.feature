Feature: Sensor value

    In order to work with a measurement
    As a developper
    I need to be able to read a sensor value

    Scenario: Read measurement from a sensor
        Given there is an empty sensor
        When i read the sensor value
        Then the value should be nil