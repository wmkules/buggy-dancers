import * as React from "react";
import axios from "axios";

import {
  addVotedIDToStorage,
  isVotedIdInStorage,
} from "../interfaces/votedCookieInterface";

import BallotInterface from "../interfaces/ballotInterface";
import VoteInterface from "../interfaces/voteInterface";
import promptInterface from "../interfaces/promptInterface";
import { BallotContext } from "../context/BallotContext";

export default function BallotComponent() {
  const { currentBallot, setCurrentBallot } = React.useContext(BallotContext);

  // does the post request for the actual vote
  const DoVote = (p: promptInterface) => {
    const vote: VoteInterface = { ballotID: currentBallot.id, promptID: p.id };
    axios
      .post<BallotInterface>("http://localhost:8080/vote", vote)
      .then((response) => {
        setCurrentBallot(response.data);
      })
      .catch((ex) => {
        const error = "An unexpected error has occurred. Could not vote";
      });
    addVotedIDToStorage(currentBallot.id);
  };

  return (
    <div>
      <h1>Vote for this</h1>
      <h2>Ballot is: {currentBallot.description}</h2>
      <h3>Prompts are:</h3>
      <ul>
        {currentBallot.prompts.map((p) => (
          <li key={p.id}>
            <p>name: {p.name}</p>
            <p>description: {p.description}</p>
            <p>votes: {p.votes}</p>
            {
              // only render the button if the ballot id is not in the list of already voted ids
              //!isVotedIdInStorage(currentBallot.id) && (
                <button onClick={() => DoVote(p)}>Vote</button>
              //)
            }
          </li>
        ))}
      </ul>
    </div>
  );
}
