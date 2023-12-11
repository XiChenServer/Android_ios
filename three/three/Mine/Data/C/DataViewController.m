//
//  DataViewController.m
//  three
//
//  Created by Huc on 2023/12/10.
//

#import "DataViewController.h"
#import "AlterViewController.h"

@interface DataViewController ()

@end

@implementation DataViewController

- (void)loadTableView {
    
    self.tableView = [[UITableView alloc] initWithFrame:self.view.bounds style:UITableViewStyleGrouped];
    self.tableView.sectionFooterHeight = 4;
    self.tableView.sectionHeaderHeight = 4;
    self.tableView.delegate = self;
    self.tableView.dataSource = self;
    self.tableView.showsVerticalScrollIndicator = NO;
    [self.view addSubview: self.tableView];
    
    self.data = @[@[@"头像", @"昵称"], @[@"手机号", @"邮箱号", @"微信号"], @[@"我的地址", @"更改密码"]];
}

- (void)viewDidLoad {
    [super viewDidLoad];
    self.title = @"个人信息";
    [self loadTableView];
    // Do any additional setup after loading the view.
}

- (UITableViewCell *)tableView:(UITableView *)tableView cellForRowAtIndexPath:(NSIndexPath *)indexPath {
    
    UITableViewCell *cell = [[UITableViewCell alloc] init];
    
    cell.textLabel.text = self.data[indexPath.section][indexPath.row];
    [cell.contentView addSubview:cell.textLabel];
    
    cell.accessoryType = UITableViewCellAccessoryDisclosureIndicator;
    
    return cell;
}

- (void)tableView:(UITableView *)tableView didSelectRowAtIndexPath:(NSIndexPath *)indexPath {
    [tableView deselectRowAtIndexPath: indexPath animated: YES];
    AlterViewController *alter = [[AlterViewController alloc] init];
    alter.row = indexPath.section*10 + indexPath.row;
    [self.navigationController pushViewController:alter animated:YES];
    
}

- (NSInteger)numberOfSectionsInTableView:(UITableView *)tableView {
    return self.data.count;
}

- (NSInteger)tableView:(UITableView *)tableView numberOfRowsInSection:(NSInteger)section {
    NSArray *temp = self.data[section];
    return temp.count;
}

- (CGFloat)tableView:(UITableView *)tableView heightForRowAtIndexPath:(NSIndexPath *)indexPath {
    return 60;
}

/*
#pragma mark - Navigation

// In a storyboard-based application, you will often want to do a little preparation before navigation
- (void)prepareForSegue:(UIStoryboardSegue *)segue sender:(id)sender {
    // Get the new view controller using [segue destinationViewController].
    // Pass the selected object to the new view controller.
}
*/

@end
