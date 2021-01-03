use serde::{Deserialize, Serialize};
use std::error::Error;
use std::io;

//HEADER: URL,Tag,Nom,Nome,Audio

#[derive(Deserialize, Debug, Serialize)]
pub struct Asset {
    #[serde(alias = "URL")]
    url: String,
    #[serde(alias = "Tag")]
    tag: String,
    #[serde(alias = "Nom")]
    mot: String,
    #[serde(alias = "Nome")]
    palavra: String,
    #[serde(alias = "Audio")]
    sound: String,
}

impl std::fmt::Display for Asset {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(
            f,
            "URL: {url}\nPalavra:\n\tPortuguês: {palavra}\n\tFrancês: {mot}\n\tTags: {tag}\n\tAudio: {sound}",
            url = self.url,
            tag = self.tag,
            mot = self.mot,
            palavra = self.palavra,
            sound = self.sound
        )
    }
}

pub async fn parse() -> Result<(), Box<dyn Error>> {
    // Build the CSV reader and iterate over each record.
    let mut rdr = csv::Reader::from_reader(io::stdin());
    for result in rdr.deserialize() {
        // The iterator yields Result<StringRecord, Error>, so we check the
        // error here.
        let record: Asset = result?;
        println!("{}", record);
    }
    Ok(())
}
