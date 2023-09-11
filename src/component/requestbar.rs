#[derive(Debug, Default)]
pub struct RequestBar {
    pub body: String,
    pub request_menu: RequestMenu,
}

impl RequestBar {
    pub fn left(&mut self) {
        self.request_menu = match self.request_menu {
            RequestMenu::Params => RequestMenu::Body,
            RequestMenu::Authentication => RequestMenu::Params,
            RequestMenu::Headers => RequestMenu::Authentication,
            RequestMenu::Body => RequestMenu::Headers,
        }
    }

    pub fn right(&mut self) {
        self.request_menu = match self.request_menu {
            RequestMenu::Params => RequestMenu::Authentication,
            RequestMenu::Authentication => RequestMenu::Headers,
            RequestMenu::Headers => RequestMenu::Body,
            RequestMenu::Body => RequestMenu::Params,
        }
    }
}

#[derive(Debug, Default, strum::Display, strum::EnumIter, PartialEq)]
pub enum RequestMenu {
    #[default]
    Params,
    Authentication,
    Headers,
    Body,
}
