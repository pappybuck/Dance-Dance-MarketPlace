using System.ComponentModel;
using System.ComponentModel.DataAnnotations.Schema;
using Microsoft.SqlServer.Server;

namespace Janus.Identity;

public class Profile
{
    [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
    public string ProfileId { get; set; }
    public string FirstName { get; set; }
    public string LastName { get; set; }
    public string Email { get; set; }
    public string Phone { get; set; }
    public string UserId { get; set; }
    public User User { get; set; }
}