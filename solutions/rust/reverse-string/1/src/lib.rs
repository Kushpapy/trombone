pub fn reverse(input: &str) -> String {

    if input == "" {
        return String::from(input);
    }
    
   // let text = String::from(input);
    let mut vec_chars:Vec<char> = input.chars().collect();
    
    let mut start = 0;
    let mut end = vec_chars.len() - 1;

    

    while start < end {
        let mut temp = vec_chars[start];
        vec_chars[start] = vec_chars[end];
        vec_chars[end] = temp;
        //vec_chars.swap(start, end);
        start += 1;
        end -= 1;
    }

    vec_chars.into_iter().collect()
}
