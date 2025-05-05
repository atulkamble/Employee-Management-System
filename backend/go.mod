module backend

go 1.19

require (
github.com/gorilla/mux v1.8.0
github.com/lib/pq v1.10.4
)

// ---------------------- frontend/src/App.js ----------------------
import React, { useEffect, useState } from 'react';
import axios from 'axios';

function App() {
const \[employees, setEmployees] = useState(\[]);
const \[name, setName] = useState("");

useEffect(() => {
axios.get("[http://localhost:8080/employees](http://localhost:8080/employees)")
.then(res => setEmployees(res.data));
}, \[]);

const addEmployee = () => {
axios.post("[http://localhost:8080/employees](http://localhost:8080/employees)", { name })
.then(() => {
setEmployees(\[...employees, { name }]);
setName("");
});
};

return ( <div> <h1>Employee Management System</h1>
\<input value={name} onChange={(e) => setName(e.target.value)} /> <button onClick={addEmployee}>Add</button> <ul>
{employees.map((emp, i) => ( <li key={i}>{emp.name}</li>
))} </ul> </div>
);
}

export default App;
