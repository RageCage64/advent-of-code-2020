using System;
using System.IO;
using System.Collections.Generic;
using System.Linq;

namespace csharp {
    class Program {
        static void Main(string[] args) {
            IEnumerable<long> input = ReadInput();
            long x = SolvePart1(input, 25);
            SolvePart2(input, x);
        }

        static IEnumerable<long> ReadInput() =>
            File.ReadAllLines("../input.txt").Select(line => Int64.Parse(line));

        static long SolvePart1(IEnumerable<long> input, long target) {
            long[] inputArr = input.ToArray();
            var preamble = new HashSet<long>();
            for(int i = 0; i < inputArr.Length; i++) {
                if(i < target) {
                    preamble.Add(inputArr[i]);
                    continue;
                }
                if(i > target) {
                    preamble.Remove(inputArr[i-(target+1)]);
                    preamble.Add(inputArr[i-1]);
                }

                bool validNumber = false;
                foreach(long num in preamble) {
                    long complement = inputArr[i] - num;
                    if(complement != num && preamble.Contains(complement)) {
                        validNumber = true;
                        break;
                    }
                }

                if(!validNumber) {
                    Console.WriteLine(inputArr[i]);
                    return inputArr[i];
                }
            }
            return 0;
        }

        static void SolvePart2(IEnumerable<long> input, long target) {
            var sequence = new List<long>();
            foreach(long num in input) {
                if(sequence.Sum() <= target) {
                    sequence.Add(num);
                }
                while(sequence.Sum() > target && sequence.Count() > 0) {
                    sequence.RemoveAt(0);
                }
                if(sequence.Sum() == target) {
                    long result = sequence.First() + sequence.Last();
                    Console.WriteLine(result);        
                    break;
                }
            }
        }
    }
}
