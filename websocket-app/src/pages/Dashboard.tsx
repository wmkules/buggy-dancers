import React from "react";
import axios from "axios";
import BallotComponent from "../components/BallotComponent";
import { BallotContext } from "../context/BallotContext";
import BallotInterface from "../interfaces/ballotInterface";

export default function Dashboard() {
    const [ballots, setBallots] = React.useState<any[]>([]);
    const fetchUsers = async () => {
        const response = await axios.get("http://localhost:8080/mysecretkey/ballots");
        setBallots(response.data);
        console.log(response.data[0].id);
        console.log("test log");
    };
    
    React.useEffect(() => {
        fetchUsers();
    }, []);

    const cueThis = (p) => {
        axios
          .get("http://localhost:8080/mysecretkey/ballots/setCurrent/" + p.id)
          .then((response) => {
            console.log(response.data);
          })
          .catch((ex) => {
            const error = "An unexpected error has occurred. Could not vote";
          });        
      };

    return (
        <div>
            {ballots.map((p) => (
          <li key={p.id}>
            <p>{p.description}<button onClick={() => cueThis(p)}>Cue</button>    </p>            
          </li>
        ))}          
        </div>
        );
}