// use crossterm::style::Stylize;
use cursorvec::CursorVec;
use ratatui::{prelude::*, style::Stylize};
use std::{
    cell::RefCell,
    collections::VecDeque,
    ops::{Deref, DerefMut},
    rc::Rc,
};

use lazy_static::lazy_static;
lazy_static! {
    static ref SELECTED_STYLE: ratatui::style::Style = Style::default().fg(Color::LightBlue);
}

#[derive(Debug, Default)]
pub struct ItemInner {
    pub name: String,
    pub selected: bool,
    pub active: bool,
    pub children: Option<ItemVec>,
}

impl ItemInner {
    pub fn new(name: &str) -> Self {
        ItemInner {
            name: name.to_string(),
            selected: false,
            active: false,
            children: None,
        }
    }
}

impl std::fmt::Display for ItemInner {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        write!(f, "{}", self.name)
    }
}

#[derive(Debug, Clone, Default)]
pub struct Item(Rc<RefCell<ItemInner>>);

impl Deref for Item {
    type Target = Rc<RefCell<ItemInner>>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for Item {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}

impl Item {
    pub fn new(name: &str) -> Self {
        Item(Rc::new(RefCell::new(ItemInner::new(name))))
    }

    pub fn with_children(self, children: Vec<Item>) -> Self {
        let mut child_refs = Vec::new();

        for child in children {
            child_refs.push(child);
        }
        self.borrow_mut().children = Some(ItemVec(CursorVec::new().with_container(child_refs)));
        self
    }

    pub fn print_item(&self, depth: usize) -> String {
        let item = self.borrow();
        let indent = "  ".repeat(depth);

        let child_str = item
            .children
            .as_ref()
            .map(|children| children.print_item(depth + 1))
            .unwrap_or_default();

        let item_str = match item.selected {
            false => item.to_string(),
            true => format!("*{}", item.to_string()),
        };

        format!("{}{}\n{}", indent, item_str, child_str)
    }

  pub fn to_span(&self) -> Span {
    let item = self.borrow();
    match item.selected {
      false => Span::raw(item.to_string()),
      true => Span::styled(item.to_string(), *SELECTED_STYLE)
    }
  }

}

impl std::fmt::Display for Item {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        let item = self.borrow();
        write!(f, "{}", item.name)
    }
}

#[derive(Debug, Default)]
pub struct ItemVec(CursorVec<Item>);

impl Deref for ItemVec {
    type Target = CursorVec<Item>;

    fn deref(&self) -> &Self::Target {
        &self.0
    }
}

impl DerefMut for ItemVec {
    fn deref_mut(&mut self) -> &mut Self::Target {
        &mut self.0
    }
}

impl ItemVec {
    pub fn new(inner: CursorVec<Item>) -> Self {
        ItemVec(inner)
    }

    pub fn preorder_iter(&self) -> Self {
        let mut stack = VecDeque::new();
        let mut result = Vec::new();

        for item in self.0.iter() {
            stack.push_back(item.clone());
        }

        while let Some(item) = stack.pop_front() {
            result.push(item.clone());
            if let Some(children) = &item.borrow().children {
                let childs = children.preorder_iter();
                for child in childs.0.iter().rev() {
                    // push children in reverse so left-most are processed first
                    result.push(child.clone());
                }
            }
        }

        ItemVec(CursorVec::new().with_container(result))
    }

    pub fn print_item(&self, depth: usize) -> String {
        self.iter().map(|item| item.print_item(depth)).collect()
    }
}
