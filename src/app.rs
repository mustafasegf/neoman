use std::error;

use cursorvec::CursorVec;

use crate::items::{Item, ItemVec};

/// Application result type.
pub type AppResult<T> = std::result::Result<T, Box<dyn error::Error>>;

/// Application.
#[derive(Debug)]
pub struct App {
    pub running: bool,
    pub counter: u8,
    pub sidebar: SideBar,
    pub settings: Settings,
}

#[derive(Debug)]
pub struct SideBar {
    pub size: u16,
    pub selected: usize,
    pub items: ItemVec,
    pub list: ItemVec,
}

#[derive(Debug)]
pub struct Settings {
    pub show_sidebar: bool,
    pub show_help: bool,
}

impl Default for App {
    fn default() -> Self {
        let items = ItemVec::new(CursorVec::new().with_container(vec![
            Item::new("Folder 1").with_children(vec![
                Item::new("Item 1.1"),
                Item::new("Item 1.2"),
                Item::new("Folder 1.3")
                    .with_children(vec![Item::new("Item 1.3.1"), Item::new("Item 1.3.2")]),
                Item::new("Item 1.4"),
            ]),
            Item::new("Item 2"),
        ]));

        let mut list = items.preorder_iter();
        list.set_rotatable(true);
        list.get_current().value().unwrap().borrow_mut().selected = true;
        items.get_current().value().unwrap().borrow_mut().selected = true;

        Self {
            running: true,
            counter: 0,
            sidebar: SideBar {
                size: 25,
                selected: 0,
                list,
                items,
            },
            settings: Settings {
                show_sidebar: true,
                show_help: false,
            },
        }
    }
}

impl App {
    /// Constructs a new instance of [`App`].
    pub fn new() -> Self {
        Self::default()
    }

    /// Handles the tick event of the terminal.
    pub fn tick(&self) {}

    /// Set running to false to quit the application.
    pub fn quit(&mut self) {
        self.running = false;
    }

    pub fn increment_counter(&mut self) {
        if let Some(res) = self.counter.checked_add(1) {
            self.counter = res;
        }
    }

    pub fn decrement_counter(&mut self) {
        if let Some(res) = self.counter.checked_sub(1) {
            self.counter = res;
        }
    }

    pub fn toggle_sidebar(&mut self) {
        self.settings.show_sidebar = !self.settings.show_sidebar;
    }

    pub fn sidebar_size(&self) -> u16 {
        match self.settings.show_sidebar {
            true => self.sidebar.size,
            false => 0,
        }
    }
}
