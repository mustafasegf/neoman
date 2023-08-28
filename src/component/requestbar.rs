#[derive(Debug, Default)]
pub struct RequestBar {
    pub body: String,
    pub request_menu: RequestMenu,
}


#[derive(Debug, Default, strum::Display, strum::EnumIter)]
pub enum RequestMenu {
    #[default]
    Params,
    Authentication,
    Headers,
    Body,
}

