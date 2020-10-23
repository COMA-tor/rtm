Feature: Receive data and put it in log file

  Scenario: Receive data and write it
    Given There is a log consumer
    When The log consumer is running
    Then It should write received data in file

  Scenario: Receive only one data
    Given There is a log consumer
    When The log consumer is running
    And The log consumer receive 1 slice of bytes
    Then 1 line should be written in file