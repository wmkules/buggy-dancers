import * as React from "react";
import BallotInterface from "../interfaces/ballotInterface";

export default class BallotComponent extends React.Component<BallotInterface, {}> {
  constructor(props: BallotInterface) {
    super(props);
  }
  render() {
      return (
        <div>
          <h1>Ballot Component</h1>
          <h2>Ballot is: {this.props.description}</h2>
          <h3>Prompts are:</h3>
          {this.props.prompts.map((p) => <>
            <p>name: {p.name}</p>
            <p>description: {p.description}</p>
            <p>votes: {p.votes}</p>
          </>)}
        </div>
      );
  }
}