import PromptInterface from "./promptInterface";

export default interface BallotInterface {
  id: string;
  description: string;
  prompts: PromptInterface[];
}
