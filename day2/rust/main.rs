struct Policy {
    min: usize,
    max: usize,
    character: char
}

struct Password {
    policy: Policy,
    password: String
}

impl Password {
    fn verify(&self) -> bool {
        let occurrences = self.password.matches(self.policy.character).count();
        if self.policy.min <= occurrences && occurrences <= self.policy.max {
            return true;
        }
        return false;
    }
}

fn buildPasswordStruct(password: str) -> Password {
    return Password {
        policy: Policy {
            min: 1,
            max: 1,
            character: 'a'
        },
        password: String::from("")
    };
}

fn readInput() -> Vec<Password> {
    let input = std::fs::read_to_string("./input.txt")
        .expect("File isn't there idiot!!!!!");
    let inputStr: String = String::from(input);
    let inputArr: Vec<&str> = inputStr.split('\n').collect();

    let mut passwords: Vec<Password> = Vec::new();
    passwords = inputArr.iter().map(|passwordStr| buildPasswordStruct(**passwordStr)).collect::Vec<Password>();
    passwords
}

fn main() {
    println!("Hello world");
}
