Feature: Collect measurements

    In order to collect measurements
    As a technician
    I need to be able to run measurement agents

    Scenario: Collect measurements from a single sensor
        Given that there is a sensor
        When I run an agent that use it
        Then measurements should be collected