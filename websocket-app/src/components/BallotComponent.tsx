import * as React from "react";
import axios from "axios";

import BallotInterface from "../interfaces/ballotInterface";
import VoteInterface from "../interfaces/voteInterface";
import promptInterface from "../interfaces/promptInterface";

import { BallotContext } from "../context/BallotContext";


export default function BallotComponent() {
  const { currentBallot, setCurrentBallot } = React.useContext(BallotContext);

  const DoVote = (p: promptInterface, b: BallotInterface) => {
    const vote: VoteInterface = { ballotID: b.id, promptID: p.id }
      axios
        .post<BallotInterface>('http://localhost:8080/vote', vote)
        .then((response) => {
          setCurrentBallot(response.data)
        })
        .catch(ex => {
          const error = "An unexpected error has occurred. Could not vote";
        });
  };



  return (
    <div>
      <h1>Vote for this</h1>
      <h2>Ballot is: {currentBallot.description}</h2>
      <h3>Prompts are:</h3>
      <ul>
        {currentBallot.prompts.map((p) =>
          <li key={p.id}>
            <p>name: {p.name}</p>
            <p>description: {p.description}</p>
            <p>votes: {p.votes}</p>
            <button onClick={() => DoVote(p, currentBallot)}>Vote</button>
          </li>
        )}
      </ul>
    </div>
  );

}

// export default BallotComponent