import React, { useEffect, useState } from 'react';
import axios from 'axios';

function EmployeeList() {
  const [employees, setEmployees] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/employees")
      .then((res) => setEmployees(res.data))
      .catch((err) => console.error(err));
  }, []);

  return (
    <div>
      <ul>
        {employees.map(emp => (
          <li key={emp.id}>{emp.name} - {emp.role}</li>
        ))}
      </ul>
    </div>
  );
}

export default EmployeeList;
