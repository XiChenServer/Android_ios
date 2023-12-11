//
//  MineViewController.h
//  three
//
//  Created by Huc on 2023/12/10.
//

#import <UIKit/UIKit.h>

NS_ASSUME_NONNULL_BEGIN

@interface MineViewController : UIViewController
<UITableViewDelegate, UITableViewDataSource>

@property UIImageView *backGroundView;
@property UITableView *tableView;
@property NSArray *data;
@property NSString *TOKEN;

@end

NS_ASSUME_NONNULL_END
