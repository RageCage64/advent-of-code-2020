open System
open System.IO

let strToInt str = int (str)

let getLines (filePath: string) = seq { yield! File.ReadLines filePath }

let getNumbers lines = seq {
    for line in lines do
        yield strToInt line
}

let complement sum num = sum - num

let q1 = 
    let lines = getLines "../input.txt"
    let numbers = getNumbers lines
    let numberSet = Set.ofSeq numbers
    let results = seq {
        for number in numbers do
            let c = complement 2020 number
            if numberSet.Contains c then
                yield number * c
    }
    Seq.last results

let q2 =
    let lines = getLines "../input.txt"
    let numbers = getNumbers lines
    let numberSet = Set.ofSeq numbers
    let results = seq {
        for number in numbers do
            let index = numbers |> Seq.findIndex (fun x -> x = number) 
            let c1 = complement 2020 number
            for number2 in numbers |> Seq.skip index do
                let c2 = complement c1 number2
                if numberSet.Contains c2 then
                    yield c2 * number * number2
    }
    match Seq.isEmpty results with
    | true -> -1
    | false -> Seq.last results

[<EntryPoint>]
let main argv =
    printfn "%i" q1
    printfn "%i" q2
    0