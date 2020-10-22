Feature: Receive data and put it in local file

  Scenario: Receive data and write it
    Given There is a local consumer
    When The local consumer is running
    Then It should write received data in file

  Scenario: Receive only one data
    Given There is a local consumer
    When The local consumer is running
    And The local consumer receive 1 slice of bytes
    Then It should 1 line be written in file