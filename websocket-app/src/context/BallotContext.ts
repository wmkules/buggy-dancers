import React, { createContext } from "react";
import BallotInterface from "../interfaces/ballotInterface";


const defaultBallot: BallotInterface = { description: "err: Default ballot", id: "wrong", prompts: [] };
// const [currentBallot, setCurrentBallot]: [BallotInterface, (ballots: BallotInterface) => void] = React.useState(defaultBallot);

export const BallotContext = createContext<any>({});