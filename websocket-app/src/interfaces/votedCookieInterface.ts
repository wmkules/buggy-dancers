export default interface votedIdsInterface {
    ids: string[];
}
const votedKey = "voted"

export function addVotedIDToStorage(id: string): void {
    let votedString = window.localStorage.getItem(votedKey)
    if (votedString == null) {
        votedString = ""
    }

    const votedIds: string[] = votedString.split(',')

    votedIds.push(id)

    window.localStorage.setItem(votedKey, votedIds.join(','))
}

export function isVotedIdInStorage(id: string): boolean {
    console.log("Check id: ", id, " in storage")
    let votedString = window.localStorage.getItem(votedKey)
    if (votedString == null) {
        return false
    }
    console.log('all the storage is ', votedString)
    const votedIds: string[] = votedString.split(',')

    return votedIds.includes(id)
}