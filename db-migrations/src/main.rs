use dotenv::dotenv;
use std::process;

use catalog_csv_reader::parse;

//HEADER: URL,Tag,Nom,Nome,Audio


#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    dotenv().ok();
    
    if let Err(err) = parse().await {
        println!("error running example: {}", err);
        process::exit(1);
    }
    Ok(())
}
