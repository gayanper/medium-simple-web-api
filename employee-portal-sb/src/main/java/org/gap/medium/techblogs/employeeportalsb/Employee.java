package org.gap.medium.techblogs.employeeportalsb;

import java.time.LocalDate;

public record Employee(String empNumber, String firstName, String lastName, LocalDate hiredOn) {
}
