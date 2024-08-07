Feature: Delete Account
  In order to delete an account
  As a user
  I want to be able to delete an account

  Scenario: Delete customer account
    Given I have a customer account
    When I delete the customer account
    Then the customer account should be deleted