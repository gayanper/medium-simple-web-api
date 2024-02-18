package org.gap.medium.techblogs.employeeportalsb;

import java.time.LocalDate;
import java.util.List;

import org.springframework.jdbc.core.simple.JdbcClient;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;


@RestController
@RequestMapping(path = "/employees")
public class Controller {
    private JdbcClient jdbcClient;

    public Controller(JdbcClient jdbcClient) {
        this.jdbcClient = jdbcClient;
    }

    @GetMapping    
    public List<Employee> findJoinedAfter(@RequestParam("after") LocalDate after) {
        return jdbcClient.sql("""
                SELECT emp_no AS empNumber, first_name AS firstName, last_name AS lastName, hire_date AS hiredOn  
                    FROM `employees` 
                    WHERE hire_date > :after
                """).param("after", after.toString()).query(Employee.class).list();
    }
}
