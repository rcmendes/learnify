mod quiz_loader;

use std::path::Path;

use quiz_loader::Quiz;
use url::ParseError::InvalidDomainCharacter;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // let url = "https://jsonplaceholder.typicode.com/photos";
    // let url = "https://cdn.pixabay.com/photo/2017/10/11/17/09/read-2841722_960_720.jpg";

    let csv_file = std::fs::File::open("./quizzes_202012121858.csv").unwrap();

    let quizzes = quiz_loader::parse(csv_file).unwrap();
    println!("List of quizzes");
    for quiz in quizzes.into_iter() {
        // println!("\n{}", quiz);
        save_quiz_image(&quiz).await?;
    }

    Ok(())
}

async fn save_quiz_image(quiz: &Quiz) -> Result<(), Box<dyn std::error::Error>> {
    if let Ok(filename) = extract_filename(&quiz.url) {
        let extension = extract_filename_extension(&filename);
        // println!("Filename: {}", filename);
        let filename = match extension {
            Some(ext) => {
                format!("{}.{}", quiz.palavra, ext)
            }
            _ => {
                quiz.palavra.clone()
            }
        };

        let root_folder = Path::new("../images");
        std::fs::create_dir_all(root_folder)?;
        let path = root_folder.join(&filename);
        // println!("File path: {:?}", path);

        let mut file = tokio::fs::File::create(path).await?;
        file_downloader::save(&quiz.url, &mut file).await?;

        let quiz_csv = format!(
            "\"{id}\",\"{uuid}\",\"{tag}\",\"{palavra}\",\"{mot}\",\"{image_filename}\",\"{audio_filename}\"",
            id = quiz.id,
            tag = quiz.tag,
            uuid = quiz.uuid,
            palavra = quiz.palavra,
            mot = quiz.mot,
            audio_filename = quiz.sound,
            image_filename = &filename
        );
        println!("{}", quiz_csv)
    }
    Ok(())
}

async fn _save_quiz_image_old(quiz: &Quiz) -> Result<(), Box<dyn std::error::Error>> {
    if let Ok(mut filename) = extract_filename(&quiz.url) {
        println!("Filename: {}", filename);
        let extension = Path::new(&filename)
            .extension()
            .and_then(std::ffi::OsStr::to_str);
        println!("Extension: {}", extension.unwrap_or("NONE"));

        filename = match extension {
            Some(_) => filename,
            None => quiz.palavra.clone(),
        };

        let root_folder = Path::new("../images");
        std::fs::create_dir_all(root_folder)?;
        let path = root_folder.join(filename);
        println!("File path: {:?}", path);

        let mut file = tokio::fs::File::create(path).await?;
        file_downloader::save(&quiz.url, &mut file).await?;
    }
    Ok(())
}

fn extract_filename(url: &str) -> Result<String, Box<dyn std::error::Error>> {
    let uri = url::Url::parse(url)?;
    let path = uri.path();
    let path_elements = path.split("/");
    match path_elements.last() {
        Some(filename) => Ok(filename.to_string()),
        None => Err(Box::new(InvalidDomainCharacter)),
    }
}

fn extract_filename_extension(filename: &str) -> Option<&str> {
    Path::new(filename)
        .extension()
        .and_then(std::ffi::OsStr::to_str)
}
