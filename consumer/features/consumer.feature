Feature: Receive measurements

  Scenario: Receive data from source
    Given There is a consumer
    When The consumer is running
    Then It should handle data
