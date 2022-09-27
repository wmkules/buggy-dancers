/**Data format to hold the already ID of all the ballots that the user has already voted for
 * Can use to ensure that the user does not vote twice on the same ballot
 */
export default interface votedIdsInterface {
  ids: string[];
}
const votedKey = "voted";

/** Mark the ID provided as a ballot that has already been voted on by the user*/
export function addVotedIDToStorage(id: string): void {
  let votedString = window.localStorage.getItem(votedKey);
  if (votedString == null) {
    votedString = "";
  }

  const votedIds: string[] = votedString.split(",");

  votedIds.push(id);

  window.localStorage.setItem(votedKey, votedIds.join(","));
}

/** Check whether the ID provided has already been voted on by the user */
export function isVotedIdInStorage(id: string): boolean {
  console.log("Check id: ", id, " in storage");
  let votedString = window.localStorage.getItem(votedKey);
  if (votedString == null) {
    return false;
  }
  console.log("all the storage is ", votedString);
  const votedIds: string[] = votedString.split(",");

  return votedIds.includes(id);
}
