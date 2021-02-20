use serde::{Deserialize, Serialize};
use std::error::Error;
use std::fs::File;

#[derive(Deserialize, Debug, Serialize)]
pub struct Quiz {
    #[serde(alias = "id")]
    pub id: String,
    #[serde(alias = "uuid")]
    pub uuid: String,
    #[serde(alias = "url")]
    pub url: String,
    #[serde(alias = "tag")]
    pub tag: String,
    #[serde(alias = "mot")]
    pub mot: String,
    #[serde(alias = "palavra")]
    pub palavra: String,
    #[serde(alias = "audio")]
    pub sound: String,
}

impl std::fmt::Display for Quiz {
    fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
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

pub fn parse(csv_file: File) -> Result<Vec<Quiz>, Box<dyn Error>> {
    let assets = csv::Reader::from_reader(csv_file)
        .deserialize()
        .map(|record| record.unwrap())
        .collect::<Vec<Quiz>>();
    Ok(assets)
}
