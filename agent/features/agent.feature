Feature: Collect measurements

    In order to collect measurements
    As a technician
    I need to be able to run measurement agents

    Scenario: Collect measurement regularly
        Given there is a sensor
        And there is a measurement interval of 10 milliseconds
        And there is an agent that use it
        When I run the agent for 45 milliseconds
        Then there should be 4 measurements collected